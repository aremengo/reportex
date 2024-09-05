package main

const defaultReportTemplate = `\documentclass[a4paper]{report}
\usepackage{graphicx}
\usepackage{pgfkeys}
\usepackage[a4paper, margin=2pt]{geometry}
\usepackage{color}
\usepackage{babel}
\usepackage{fontspec}
\defaultfontfeatures{Mapping=tex-text,Scale=MatchLowercase}
\usepackage[sfdefault]{roboto}
\usepackage{lipsum}  

\graphicspath{{../images/}}

\begin{document}

\begin{center}
\LARGE{Example}
\\
\includegraphics[scale=.3]{avatar.png}
\bigskip
\end{center}
\lipsum[2-4]
\end{document}
`
