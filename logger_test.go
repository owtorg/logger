package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

//func TestFmtLog(t *testing.T) {
//	fmtLog := new(FmtLog)
//	testLogLevels(fmtLog, t)
//}

func TestStdLog(t *testing.T) {
	stdLog := new(StdLog)
	output := captureOutput(func() {
		stdLog.Log("Emergency", "This is a Log message")
	})
	testOutput(output, "Emergency [This is a Log message]\n", t)

	output = captureOutput(func() {
		stdLog.Log("Arbitrary", "This is a generic message")
	})
	testOutput(output, "Arbitrary [This is a generic message]\n", t)

	testLogLevels(stdLog, t)
}

func TestStack(t *testing.T) {

	stack := new(Stack)
	stack.Add(new(FmtLog))
	stack.Add(new(StdLog))

	//This only captures the logger output, not the fmt
	//TODO - test the fmt as well
	testLogLevels(stack, t)
	output := captureOutput(func() {
		stack.Log("custom level", "Log to a custom level")
	})
	testOutput(output, "custom level [Log to a custom level]\n", t)

}

//tlCallback is passed to the fileloggers init function so that the filelog save location can be set
func tlCallback(s *FileLog) {
	filename := "./test/output/testfile" + time.Now().String()
	s.logPath = filename
}

func TestFileLog(t *testing.T) {

	//Test with custom log location set in init closure
	fl := new(FileLog)
	fl.OnInit(tlCallback)
	err := fl.Init()
	if err != nil {
		t.Error("Init failed", err)
		t.FailNow()
	}
	fl.Emergency("This is an Emergency message")
	fl.Alert("This is an Alert message")
	fl.Critical("This is a Critical message")
	fl.Error("This is an Error message")
	fl.Warning("This is a Warning message")
	fl.Notice("This is a Notice")
	fl.Info("This is an Info message")
	fl.Debug("This is a Debug message")

	b, err := ioutil.ReadFile(fl.logPath)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	if str != "Emergency [This is an Emergency message]\nAlert [This is an Alert message]\nCritical [This is a Critical message]\nError [This is an Error message]\nWarning [This is a Warning message]\nNotice [This is a Notice]\nInfo [This is an Info message]\nDebug [This is a Debug message]\n" {
		t.Error("match failed for str", str)
	}

	//Test with default location on a stack
	stack := new(Stack)
	fl2 := new(FileLog)
	stack.Add(fl2)
	stack.Emergency("This is an Emergency message")
	stack.Alert("This is an Alert message")
	stack.Critical("This is a Critical message")
	stack.Error("This is an Error message")
	stack.Warning("This is a Warning message")
	stack.Notice("This is a Notice")
	stack.Info("This is an Info message")
	stack.Debug("This is a Debug message")

	c, err := ioutil.ReadFile(fl2.logPath)
	if err != nil {
		fmt.Print(err)
	}

	str = string(c)

	if str != "Emergency [This is an Emergency message]\nAlert [This is an Alert message]\nCritical [This is a Critical message]\nError [This is an Error message]\nWarning [This is a Warning message]\nNotice [This is a Notice]\nInfo [This is an Info message]\nDebug [This is a Debug message]\n" {
		t.Error("match failed for str", str)
	}
	os.Remove(fl2.logPath)

}

func testLogLevels(stdLog Logger, t *testing.T) {
	output := captureOutput(func() {
		stdLog.Emergency("This is a message")
	})
	testOutput(output, "Emergency [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Alert("This is a message")
	})
	testOutput(output, "Alert [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Critical("This is a message")
	})
	testOutput(output, "Critical [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Error("This is a message")
	})
	testOutput(output, "Error [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Warning("This is a message")
	})
	testOutput(output, "Warning [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Notice("This is a message")
	})
	testOutput(output, "Notice [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Info("This is a message")
	})
	testOutput(output, "Info [This is a message]\n", t)

	output = captureOutput(func() {
		stdLog.Debug("This is a message")
	})
	testOutput(output, "Debug [This is a message]\n", t)
}

func testOutput(output string, expected string, t *testing.T) {
	if output != expected {
		t.Error("UNEXPECTED OUTPUT", "Expected:", expected, "Output:", output)
	}

}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
