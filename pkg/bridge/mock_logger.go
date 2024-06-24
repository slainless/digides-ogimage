package bridge

type MockLogger struct{}

func (l *MockLogger) Log(args ...interface{})   {}
func (l *MockLogger) Debug(args ...interface{}) {}
func (l *MockLogger) Info(args ...interface{})  {}
func (l *MockLogger) Warn(args ...interface{})  {}
func (l *MockLogger) Error(args ...interface{}) {}
func (l *MockLogger) Dir(data interface{})      {}
func (l *MockLogger) DirXML(data interface{})   {}

var mockLogger = &MockLogger{}
