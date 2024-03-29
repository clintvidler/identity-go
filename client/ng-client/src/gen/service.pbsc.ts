/* tslint:disable */
/* eslint-disable */
// @ts-nocheck
//
// THIS IS A GENERATED FILE
// DO NOT MODIFY IT! YOUR CHANGES WILL BE LOST
import { Inject, Injectable, Optional } from '@angular/core';
import {
  GrpcCallType,
  GrpcClient,
  GrpcClientFactory,
  GrpcEvent,
  GrpcMetadata
} from '@ngx-grpc/common';
import {
  GRPC_CLIENT_FACTORY,
  GrpcHandler,
  takeMessages,
  throwStatusErrors
} from '@ngx-grpc/core';
import { Observable } from 'rxjs';
import * as thisProto from './service.pb';
import { GRPC_IDENTITY_SERVICE_CLIENT_SETTINGS } from './service.pbconf';
/**
 * Service client implementation for proto.IdentityService
 */
@Injectable({ providedIn: 'any' })
export class IdentityServiceClient {
  private client: GrpcClient<any>;

  /**
   * Raw RPC implementation for each service client method.
   * The raw methods provide more control on the incoming data and events. E.g. they can be useful to read status `OK` metadata.
   * Attention: these methods do not throw errors when non-zero status codes are received.
   */
  $raw = {
    /**
     * Unary call: /proto.IdentityService/CurrentUser
     *
     * @param requestMessage Request message
     * @param requestMetadata Request metadata
     * @returns Observable<GrpcEvent<thisProto.CurrentUserReponse>>
     */
    currentUser: (
      requestData: thisProto.CurrentUserRequest,
      requestMetadata = new GrpcMetadata()
    ): Observable<GrpcEvent<thisProto.CurrentUserReponse>> => {
      return this.handler.handle({
        type: GrpcCallType.unary,
        client: this.client,
        path: '/proto.IdentityService/CurrentUser',
        requestData,
        requestMetadata,
        requestClass: thisProto.CurrentUserRequest,
        responseClass: thisProto.CurrentUserReponse
      });
    },
    /**
     * Unary call: /proto.IdentityService/Login
     *
     * @param requestMessage Request message
     * @param requestMetadata Request metadata
     * @returns Observable<GrpcEvent<thisProto.LoginReponse>>
     */
    login: (
      requestData: thisProto.LoginRequest,
      requestMetadata = new GrpcMetadata()
    ): Observable<GrpcEvent<thisProto.LoginReponse>> => {
      return this.handler.handle({
        type: GrpcCallType.unary,
        client: this.client,
        path: '/proto.IdentityService/Login',
        requestData,
        requestMetadata,
        requestClass: thisProto.LoginRequest,
        responseClass: thisProto.LoginReponse
      });
    },
    /**
     * Unary call: /proto.IdentityService/Logout
     *
     * @param requestMessage Request message
     * @param requestMetadata Request metadata
     * @returns Observable<GrpcEvent<thisProto.LogoutReponse>>
     */
    logout: (
      requestData: thisProto.LogoutRequest,
      requestMetadata = new GrpcMetadata()
    ): Observable<GrpcEvent<thisProto.LogoutReponse>> => {
      return this.handler.handle({
        type: GrpcCallType.unary,
        client: this.client,
        path: '/proto.IdentityService/Logout',
        requestData,
        requestMetadata,
        requestClass: thisProto.LogoutRequest,
        responseClass: thisProto.LogoutReponse
      });
    },
    /**
     * Unary call: /proto.IdentityService/Refresh
     *
     * @param requestMessage Request message
     * @param requestMetadata Request metadata
     * @returns Observable<GrpcEvent<thisProto.RefreshReponse>>
     */
    refresh: (
      requestData: thisProto.RefreshRequest,
      requestMetadata = new GrpcMetadata()
    ): Observable<GrpcEvent<thisProto.RefreshReponse>> => {
      return this.handler.handle({
        type: GrpcCallType.unary,
        client: this.client,
        path: '/proto.IdentityService/Refresh',
        requestData,
        requestMetadata,
        requestClass: thisProto.RefreshRequest,
        responseClass: thisProto.RefreshReponse
      });
    }
  };

  constructor(
    @Optional() @Inject(GRPC_IDENTITY_SERVICE_CLIENT_SETTINGS) settings: any,
    @Inject(GRPC_CLIENT_FACTORY) clientFactory: GrpcClientFactory<any>,
    private handler: GrpcHandler
  ) {
    this.client = clientFactory.createClient('proto.IdentityService', settings);
  }

  /**
   * Unary call @/proto.IdentityService/CurrentUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.CurrentUserReponse>
   */
  currentUser(
    requestData: thisProto.CurrentUserRequest,
    requestMetadata = new GrpcMetadata()
  ): Observable<thisProto.CurrentUserReponse> {
    return this.$raw
      .currentUser(requestData, requestMetadata)
      .pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Unary call @/proto.IdentityService/Login
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.LoginReponse>
   */
  login(
    requestData: thisProto.LoginRequest,
    requestMetadata = new GrpcMetadata()
  ): Observable<thisProto.LoginReponse> {
    return this.$raw
      .login(requestData, requestMetadata)
      .pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Unary call @/proto.IdentityService/Logout
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.LogoutReponse>
   */
  logout(
    requestData: thisProto.LogoutRequest,
    requestMetadata = new GrpcMetadata()
  ): Observable<thisProto.LogoutReponse> {
    return this.$raw
      .logout(requestData, requestMetadata)
      .pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Unary call @/proto.IdentityService/Refresh
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.RefreshReponse>
   */
  refresh(
    requestData: thisProto.RefreshRequest,
    requestMetadata = new GrpcMetadata()
  ): Observable<thisProto.RefreshReponse> {
    return this.$raw
      .refresh(requestData, requestMetadata)
      .pipe(throwStatusErrors(), takeMessages());
  }
}
