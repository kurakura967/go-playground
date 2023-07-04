import grpc
import hello_pb2
import hello_pb2_grpc


def run():
    print("please enter your name")
    msg = input()
    with grpc.insecure_channel('localhost:8080') as channel:
        stub = hello_pb2_grpc.GreetingServiceStub(channel)
        response = stub.Hello(hello_pb2.HelloRequest(name=msg))
        print(response.message)


if __name__ == '__main__':
    print("start gRPC Client")
    run()
