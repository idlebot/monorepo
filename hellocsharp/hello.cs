using System;
using gaxgrpc = Google.Api.Gax.Grpc;
using GreeterV1 = HelloGrpc.Greeter.V1;

namespace Hello
{
    public static class Program
    {
        public static void Main()
        {
            AppContext.SetSwitch("System.Net.Http.SocketsHttpHandler.Http2UnencryptedSupport", true);
            var channel = global::Grpc.Net.Client.GrpcChannel.ForAddress("http://localhost:5432");
            var builder = new GreeterV1.GreeterClientBuilder
            {
                CallInvoker = channel.CreateCallInvoker(),
            };
            var client = builder.Build();

            var req = new GreeterV1.HelloRequest
            {
                Name = "C#"
            };

            var response = client.Hello(req);
            Console.WriteLine(response.Message);
        }
    }
}