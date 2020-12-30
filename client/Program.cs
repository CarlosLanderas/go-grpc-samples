using System;
using System.IO;
using System.Net.Http;
using System.Security.Cryptography.X509Certificates;
using System.Threading.Channels;
using System.Threading.Tasks;
using Grpc;
using Grpc.Core;
using Grpc.Net.Client;
using static Grpc.UserService;

namespace client
{
    class Program
    {
        static async Task Main(string[] args)
        {

            var httpHandler = new HttpClientHandler();
            httpHandler.ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator;
            
            X509Certificate2 cert = X509Certificate2.CreateFromPemFile("../../../../certs/server_cert.pem",  "../../../../certs/server_key.pem");
            httpHandler.ClientCertificates.Add(cert);


            var credentials = CallCredentials.FromInterceptor((context, metadata) =>
            {
                metadata.Add("Authorization", "Bearer landetoken");
                return Task.CompletedTask;
            });

            var channelOptions = new GrpcChannelOptions
            {
                HttpHandler = httpHandler,
                Credentials = ChannelCredentials.Create(new SslCredentials(), credentials)
            };

            using var channel = GrpcChannel.ForAddress("https://localhost:8000", channelOptions);
            var client = new UserServiceClient(channel);
            var request = new GetUsersRequest();
            request.Ids.Add(3);
            request.Ids.Add(6);
            request.Ids.Add(20);

            var users = await client.GetUsersAsync(request);
        }
    }
}
