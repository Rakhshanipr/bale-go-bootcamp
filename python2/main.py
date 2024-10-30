import this
import grpc
from python2.pythonp2 import messaging_pb2_grpc, messaging_pb2

def main():
    print("Hello World!")
    channel = grpc.insecure_channel('localhost:4042')
    stub = messaging_pb2_grpc.MessagingStub(channel)
    a = stub.AddUser(messaging_pb2.RequestAddUser(user_name="moh", file_id="neda"), wait_for_ready=True)
    print("Response received:", a.Id)

if __name__ == '__main__':
    main()
