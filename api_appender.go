package logapi

// Appender appender struct
type Appender interface {
	GetAppenderName() string

	// Initialize appender
	Initialize()

	// set the layout
	SetLayout(layout Layout)

	// get open layout
	GetLayout() Layout
}
