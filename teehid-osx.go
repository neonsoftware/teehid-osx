package teensyrawhid

/*
#cgo CFLAGS: -I/usr/local/include -Wall -O2 -DOS_MACOSX -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.10.sdk -mmacosx-version-min=10.5
#cgo LDFLAGS: -mmacosx-version-min=10.5  -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.10.sdk -framework IOKit -framework CoreFoundation
#include "hid.h"
#include "hid_MACOSX.h"
 */
import "C"
import "fmt"
//import "log"
import "unsafe"

type TeensyRawHidDevice struct {
  VendorId int
  ProductId int  
  ReadDebugLevel int
  WriteDebugLevel int
  device_number int
}

func (device *TeensyRawHidDevice) Open( vendorId int, productId int ) bool {
  // C-based example is 16C0:0480:FFAB:0200
  device.VendorId = vendorId
  device.ProductId = productId
  device.device_number = 0
  opened := int( C.rawhid_open(1, C.int(device.VendorId), C.int(device.ProductId), 0xFFAB, 0x0200) );
  fmt.Println("Rawhid device opened :", opened);
  return int(opened) != 0
}

func (device *TeensyRawHidDevice)  Close(){
  C.rawhid_close(C.int(device.device_number))
}

func (device *TeensyRawHidDevice) Recv( buf_len int, timeout int ) (int, []byte) {
  buf := make([]byte, buf_len)
  recv_bytes := C.rawhid_recv(C.int(device.device_number), unsafe.Pointer(&buf[0]), C.int(len(buf)), C.int(timeout));
  //fmt.Printf(">>>>>>>> Received: %v bytes: %s\n", int(recv_bytes), string(buf));
  return int(recv_bytes), buf
}

func (device *TeensyRawHidDevice) Send( buf_len int, timeout int ) int {
  buf := make([]byte, buf_len)
  sent_bytes := C.rawhid_send(C.int(device.device_number), unsafe.Pointer(&buf[0]), C.int(len(buf)), C.int(timeout));
  //fmt.Printf(">>>>>>>> Sent: %v bytes\n", int(sent_bytes));
  return int(sent_bytes)
}
