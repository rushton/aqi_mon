import serial
import io

ser = serial.Serial("/dev/cu.usbserial-1410", 115200)


while True:
    print(ser.readline())
