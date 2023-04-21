// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated code. DO NOT EDIT!

#pragma warning disable CS8981
using gax = Google.Api.Gax;
using gaxgrpc = Google.Api.Gax.Grpc;
using proto = Google.Protobuf;
using grpccore = Grpc.Core;
using grpcinter = Grpc.Core.Interceptors;
using mel = Microsoft.Extensions.Logging;
using sys = System;
using scg = System.Collections.Generic;
using sco = System.Collections.ObjectModel;
using st = System.Threading;
using stt = System.Threading.Tasks;

namespace HelloGrpc.Greeter.V1
{
    /// <summary>Settings for <see cref="GreeterClient"/> instances.</summary>
    public sealed partial class GreeterSettings : gaxgrpc::ServiceSettingsBase
    {
        /// <summary>Get a new instance of the default <see cref="GreeterSettings"/>.</summary>
        /// <returns>A new instance of the default <see cref="GreeterSettings"/>.</returns>
        public static GreeterSettings GetDefault() => new GreeterSettings();

        /// <summary>Constructs a new <see cref="GreeterSettings"/> object with default settings.</summary>
        public GreeterSettings()
        {
        }

        private GreeterSettings(GreeterSettings existing) : base(existing)
        {
            gax::GaxPreconditions.CheckNotNull(existing, nameof(existing));
            HelloSettings = existing.HelloSettings;
            OnCopy(existing);
        }

        partial void OnCopy(GreeterSettings existing);

        /// <summary>
        /// <see cref="gaxgrpc::CallSettings"/> for synchronous and asynchronous calls to <c>GreeterClient.Hello</c> and
        /// <c>GreeterClient.HelloAsync</c>.
        /// </summary>
        /// <remarks>
        /// <list type="bullet">
        /// <item><description>This call will not be retried.</description></item>
        /// <item><description>Timeout: 600 seconds.</description></item>
        /// </list>
        /// </remarks>
        public gaxgrpc::CallSettings HelloSettings { get; set; } = gaxgrpc::CallSettings.FromExpiration(gax::Expiration.FromTimeout(sys::TimeSpan.FromMilliseconds(600000)));

        /// <summary>Creates a deep clone of this object, with all the same property values.</summary>
        /// <returns>A deep clone of this <see cref="GreeterSettings"/> object.</returns>
        public GreeterSettings Clone() => new GreeterSettings(this);
    }

    /// <summary>
    /// Builder class for <see cref="GreeterClient"/> to provide simple configuration of credentials, endpoint etc.
    /// </summary>
    public sealed partial class GreeterClientBuilder : gaxgrpc::ClientBuilderBase<GreeterClient>
    {
        /// <summary>The settings to use for RPCs, or <c>null</c> for the default settings.</summary>
        public GreeterSettings Settings { get; set; }

        /// <summary>Creates a new builder with default settings.</summary>
        public GreeterClientBuilder() : base(GreeterClient.ServiceMetadata)
        {
        }

        partial void InterceptBuild(ref GreeterClient client);

        partial void InterceptBuildAsync(st::CancellationToken cancellationToken, ref stt::Task<GreeterClient> task);

        /// <summary>Builds the resulting client.</summary>
        public override GreeterClient Build()
        {
            GreeterClient client = null;
            InterceptBuild(ref client);
            return client ?? BuildImpl();
        }

        /// <summary>Builds the resulting client asynchronously.</summary>
        public override stt::Task<GreeterClient> BuildAsync(st::CancellationToken cancellationToken = default)
        {
            stt::Task<GreeterClient> task = null;
            InterceptBuildAsync(cancellationToken, ref task);
            return task ?? BuildAsyncImpl(cancellationToken);
        }

        private GreeterClient BuildImpl()
        {
            Validate();
            grpccore::CallInvoker callInvoker = CreateCallInvoker();
            return GreeterClient.Create(callInvoker, Settings, Logger);
        }

        private async stt::Task<GreeterClient> BuildAsyncImpl(st::CancellationToken cancellationToken)
        {
            Validate();
            grpccore::CallInvoker callInvoker = await CreateCallInvokerAsync(cancellationToken).ConfigureAwait(false);
            return GreeterClient.Create(callInvoker, Settings, Logger);
        }

        /// <summary>Returns the channel pool to use when no other options are specified.</summary>
        protected override gaxgrpc::ChannelPool GetChannelPool() => GreeterClient.ChannelPool;
    }

    /// <summary>Greeter client wrapper, for convenient use.</summary>
    /// <remarks>
    /// The greeting service definition.
    /// </remarks>
    public abstract partial class GreeterClient
    {
        /// <summary>
        /// The default endpoint for the Greeter service, which is a host of "greeter.googleapis.com" and a port of 443.
        /// </summary>
        public static string DefaultEndpoint { get; } = "greeter.googleapis.com:443";

        /// <summary>The default Greeter scopes.</summary>
        /// <remarks>The default Greeter scopes are:<list type="bullet"></list></remarks>
        public static scg::IReadOnlyList<string> DefaultScopes { get; } = new sco::ReadOnlyCollection<string>(new string[] { });

        /// <summary>The service metadata associated with this client type.</summary>
        public static gaxgrpc::ServiceMetadata ServiceMetadata { get; } = new gaxgrpc::ServiceMetadata(Greeter.Descriptor, DefaultEndpoint, DefaultScopes, true, gax::ApiTransports.Grpc, PackageApiMetadata.ApiMetadata);

        internal static gaxgrpc::ChannelPool ChannelPool { get; } = new gaxgrpc::ChannelPool(ServiceMetadata);

        /// <summary>
        /// Asynchronously creates a <see cref="GreeterClient"/> using the default credentials, endpoint and settings. 
        /// To specify custom credentials or other settings, use <see cref="GreeterClientBuilder"/>.
        /// </summary>
        /// <param name="cancellationToken">
        /// The <see cref="st::CancellationToken"/> to use while creating the client.
        /// </param>
        /// <returns>The task representing the created <see cref="GreeterClient"/>.</returns>
        public static stt::Task<GreeterClient> CreateAsync(st::CancellationToken cancellationToken = default) =>
            new GreeterClientBuilder().BuildAsync(cancellationToken);

        /// <summary>
        /// Synchronously creates a <see cref="GreeterClient"/> using the default credentials, endpoint and settings. To
        /// specify custom credentials or other settings, use <see cref="GreeterClientBuilder"/>.
        /// </summary>
        /// <returns>The created <see cref="GreeterClient"/>.</returns>
        public static GreeterClient Create() => new GreeterClientBuilder().Build();

        /// <summary>
        /// Creates a <see cref="GreeterClient"/> which uses the specified call invoker for remote operations.
        /// </summary>
        /// <param name="callInvoker">
        /// The <see cref="grpccore::CallInvoker"/> for remote operations. Must not be null.
        /// </param>
        /// <param name="settings">Optional <see cref="GreeterSettings"/>.</param>
        /// <param name="logger">Optional <see cref="mel::ILogger"/>.</param>
        /// <returns>The created <see cref="GreeterClient"/>.</returns>
        internal static GreeterClient Create(grpccore::CallInvoker callInvoker, GreeterSettings settings = null, mel::ILogger logger = null)
        {
            gax::GaxPreconditions.CheckNotNull(callInvoker, nameof(callInvoker));
            grpcinter::Interceptor interceptor = settings?.Interceptor;
            if (interceptor != null)
            {
                callInvoker = grpcinter::CallInvokerExtensions.Intercept(callInvoker, interceptor);
            }
            Greeter.GreeterClient grpcClient = new Greeter.GreeterClient(callInvoker);
            return new GreeterClientImpl(grpcClient, settings, logger);
        }

        /// <summary>
        /// Shuts down any channels automatically created by <see cref="Create()"/> and
        /// <see cref="CreateAsync(st::CancellationToken)"/>. Channels which weren't automatically created are not
        /// affected.
        /// </summary>
        /// <remarks>
        /// After calling this method, further calls to <see cref="Create()"/> and
        /// <see cref="CreateAsync(st::CancellationToken)"/> will create new channels, which could in turn be shut down
        /// by another call to this method.
        /// </remarks>
        /// <returns>A task representing the asynchronous shutdown operation.</returns>
        public static stt::Task ShutdownDefaultChannelsAsync() => ChannelPool.ShutdownChannelsAsync();

        /// <summary>The underlying gRPC Greeter client</summary>
        public virtual Greeter.GreeterClient GrpcClient => throw new sys::NotImplementedException();

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request object containing all of the parameters for the API call.</param>
        /// <param name="callSettings">If not null, applies overrides to this RPC call.</param>
        /// <returns>The RPC response.</returns>
        public virtual HelloReply Hello(HelloRequest request, gaxgrpc::CallSettings callSettings = null) =>
            throw new sys::NotImplementedException();

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request object containing all of the parameters for the API call.</param>
        /// <param name="callSettings">If not null, applies overrides to this RPC call.</param>
        /// <returns>A Task containing the RPC response.</returns>
        public virtual stt::Task<HelloReply> HelloAsync(HelloRequest request, gaxgrpc::CallSettings callSettings = null) =>
            throw new sys::NotImplementedException();

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request object containing all of the parameters for the API call.</param>
        /// <param name="cancellationToken">A <see cref="st::CancellationToken"/> to use for this RPC.</param>
        /// <returns>A Task containing the RPC response.</returns>
        public virtual stt::Task<HelloReply> HelloAsync(HelloRequest request, st::CancellationToken cancellationToken) =>
            HelloAsync(request, gaxgrpc::CallSettings.FromCancellationToken(cancellationToken));
    }

    /// <summary>Greeter client wrapper implementation, for convenient use.</summary>
    /// <remarks>
    /// The greeting service definition.
    /// </remarks>
    public sealed partial class GreeterClientImpl : GreeterClient
    {
        private readonly gaxgrpc::ApiCall<HelloRequest, HelloReply> _callHello;

        /// <summary>
        /// Constructs a client wrapper for the Greeter service, with the specified gRPC client and settings.
        /// </summary>
        /// <param name="grpcClient">The underlying gRPC client.</param>
        /// <param name="settings">The base <see cref="GreeterSettings"/> used within this client.</param>
        /// <param name="logger">Optional <see cref="mel::ILogger"/> to use within this client.</param>
        public GreeterClientImpl(Greeter.GreeterClient grpcClient, GreeterSettings settings, mel::ILogger logger)
        {
            GrpcClient = grpcClient;
            GreeterSettings effectiveSettings = settings ?? GreeterSettings.GetDefault();
            gaxgrpc::ClientHelper clientHelper = new gaxgrpc::ClientHelper(effectiveSettings, logger);
            _callHello = clientHelper.BuildApiCall<HelloRequest, HelloReply>("Hello", grpcClient.HelloAsync, grpcClient.Hello, effectiveSettings.HelloSettings).WithGoogleRequestParam("name", request => request.Name);
            Modify_ApiCall(ref _callHello);
            Modify_HelloApiCall(ref _callHello);
            OnConstruction(grpcClient, effectiveSettings, clientHelper);
        }

        partial void Modify_ApiCall<TRequest, TResponse>(ref gaxgrpc::ApiCall<TRequest, TResponse> call) where TRequest : class, proto::IMessage<TRequest> where TResponse : class, proto::IMessage<TResponse>;

        partial void Modify_HelloApiCall(ref gaxgrpc::ApiCall<HelloRequest, HelloReply> call);

        partial void OnConstruction(Greeter.GreeterClient grpcClient, GreeterSettings effectiveSettings, gaxgrpc::ClientHelper clientHelper);

        /// <summary>The underlying gRPC Greeter client</summary>
        public override Greeter.GreeterClient GrpcClient { get; }

        partial void Modify_HelloRequest(ref HelloRequest request, ref gaxgrpc::CallSettings settings);

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request object containing all of the parameters for the API call.</param>
        /// <param name="callSettings">If not null, applies overrides to this RPC call.</param>
        /// <returns>The RPC response.</returns>
        public override HelloReply Hello(HelloRequest request, gaxgrpc::CallSettings callSettings = null)
        {
            Modify_HelloRequest(ref request, ref callSettings);
            return _callHello.Sync(request, callSettings);
        }

        /// <summary>
        /// Sends a greeting
        /// </summary>
        /// <param name="request">The request object containing all of the parameters for the API call.</param>
        /// <param name="callSettings">If not null, applies overrides to this RPC call.</param>
        /// <returns>A Task containing the RPC response.</returns>
        public override stt::Task<HelloReply> HelloAsync(HelloRequest request, gaxgrpc::CallSettings callSettings = null)
        {
            Modify_HelloRequest(ref request, ref callSettings);
            return _callHello.Async(request, callSettings);
        }
    }
}
