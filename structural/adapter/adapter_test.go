package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		//此处创建一个MyLegacyPrinter实例，赋给OldPrinter
		Msg: msg,
	}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
	adapter = PrinterAdapter{
		OldPrinter: nil,
		//此处传递一个空值给OldPrinter，期望按原样输出
		Msg: msg,
	}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
