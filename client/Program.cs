using System;
using System.Threading.Tasks;
using Grpc;
using Grpc.Net.Client;
using static Grpc.UserService;

namespace client
{
    class Program
    {
        static async Task Main(string[] args)
        {
            using var channel = GrpcChannel.ForAddress("http://localhost:8000");
            var client = new UserServiceClient(channel);
            var request = new GetUsersRequest();
            request.Ids.Add(3);
            request.Ids.Add(6);
            request.Ids.Add(20);

            var users = await client.GetUsersAsync(request);
        }
    }
}
