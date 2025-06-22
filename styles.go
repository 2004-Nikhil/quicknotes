package main

import "github.com/charmbracelet/lipgloss"

// Color definitions
var (
    primaryColor   = lipgloss.Color("#ff6b9d")
    secondaryColor = lipgloss.Color("#4ecdc4")
    accentColor    = lipgloss.Color("#45b7d1")
    successColor   = lipgloss.Color("#96ceb4")
    warningColor   = lipgloss.Color("#ffeaa7")
    errorColor     = lipgloss.Color("#fd79a8")
    textColor      = lipgloss.Color("#2d3436")
    subtleColor    = lipgloss.Color("#636e72")
)

// Styles
var (
    titleStyle        = lipgloss.NewStyle().Foreground(primaryColor).Bold(true).BorderStyle(lipgloss.NormalBorder()).BorderForeground(primaryColor).Padding(0, 1)
    headerStyle       = lipgloss.NewStyle().Foreground(secondaryColor).Bold(true).Margin(1, 0)
    itemStyle         = lipgloss.NewStyle().Foreground(textColor).Padding(0, 1)
    selectedItemStyle = lipgloss.NewStyle().Foreground(primaryColor).Background(lipgloss.Color("#ffffff")).Bold(true).Padding(0, 1)
    tagStyle          = lipgloss.NewStyle().Foreground(accentColor).Background(lipgloss.Color("#e8f4fd")).Padding(0, 1).Margin(0, 1)
    folderStyle       = lipgloss.NewStyle().Foreground(successColor).Bold(true)
    helpStyle         = lipgloss.NewStyle().Foreground(subtleColor).Margin(1, 0)
)