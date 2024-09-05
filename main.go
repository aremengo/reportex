package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	slog.Info("generating a new report", "id", id.String())

	tmpDir, err := os.MkdirTemp("/app", "report*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // clean up the temp directory

	slog.Debug("using temp directory", "dir", tmpDir)

	f, err := os.Create(path.Join(tmpDir, "report.tex"))
	if err != nil {
		log.Fatal("cannot create temporary tex file")
	}

	f.WriteString(defaultReportTemplate)
	f.Close()

	cmd := exec.Command("xelatex", "--interaction=nonstopmode", "--output-directory="+tmpDir, "-file-line-error", "report.tex")
	cmd.Dir = tmpDir

	out, err := cmd.CombinedOutput()
	slog.Info("calling XeLaTeX to generate the pdf")
	if err != nil {
		slog.Error("error calling XeLaTeX", "err", err, "output", string(out))
		return
	}

	// move the generate pdf out of the temporary folder

	// this doesn't work when the target directory is volume mounted
	//// err = os.Rename(path.Join(tmpDir, "report.pdf"), path.Join("/app", "reports", "report.pdf"))
	//// if err != nil {
	//// 	fmt.Println(err)
	//// }

	reportFilename := fmt.Sprintf("report-%s.pdf", id.String())

	src := path.Join(tmpDir, "report.pdf")
	dest := path.Join("/app", "reports", reportFilename)

	if err = exec.Command("mv", src, dest).Run(); err != nil {
		slog.Error("error moving the generated file", "err", err)
	}
}
