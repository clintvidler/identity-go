// import { ApplicationConfig } from '@angular/core';
// import { provideRouter } from '@angular/router';

// import { GrpcCoreModule } from '@ngx-grpc/core';
// import { GrpcWebClientModule } from '@ngx-grpc/grpc-web-client';

// import { routes } from './app.routes';

// export const appConfig: ApplicationConfig = {
//   imports: [
//     GrpcCoreModule.forRoot(),
//     GrpcWebClientModule.forRoot({
//       settings: { host: 'http://localhost:9900' }
//     })
//   ],
//   providers: [
//     provideRouter(routes),
//   ]
// };

import { ApplicationRef, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { provideRouter } from '@angular/router';

import { GrpcCoreModule } from '@ngx-grpc/core';
import {
  GrpcWebClientModule,
  GrpcWebClientSettings
} from '@ngx-grpc/grpc-web-client';

import { routes } from './app.routes';
import { AppComponent } from './app.component';

@NgModule({
  imports: [
    BrowserModule,
    GrpcCoreModule.forRoot(),
    GrpcWebClientModule.forRoot({
      settings: {
        host: 'http://localhost:9800',
        modifyHeaders: (headers: any) => {
          // Optional: Modify headers if needed
          // Add any custom headers if required
          headers['Content-Type'] = 'application/grpc-web-text';
          return headers;
        }
        // modifyUrl: (url: string) => {
        //   // Modify the URL to remove the service name prefix
        //   return url.replace('Login', 'hh');
        // }
        // transport: environment.grpcTransport // Optional: Use different transport if needed
      } as GrpcWebClientSettings // Type assertion
    })
  ],
  providers: [provideRouter(routes)]
})
export class AppConfigModule {
  constructor(private appRef: ApplicationRef) {}

  ngDoBootstrap() {
    this.appRef.bootstrap(AppComponent); // Assuming AppComponent is your main component
  }
}
