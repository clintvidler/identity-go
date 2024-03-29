import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { IdentityServiceClient } from '../gen/service.pbsc';
import { LoginRequest } from '../gen/service.pb';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  providers: [],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'ng-client';

  constructor(private client: IdentityServiceClient) {}

  // sendOne() {
  //   this.client.echo(new EchoRequest({ message: 'text' })).subscribe(res => console.log(res));

  //   // or if you want to terminate it, e.g. it is a server stream or you navigate away and do not need to wait
  //   const sub = this.client.echo(new EchoRequest({ message: 'text' })).subscribe(res => console.log(res));

  //   setTimeout(() => sub.unsubscribe(), 1000); // this closes connection
  // }

  sendOne() {
    this.client
      .login(new LoginRequest({ email: 'x@x', password: 'x' }))
      .subscribe(res => console.log(res));

    // // or if you want to terminate it, e.g. it is a server stream or you navigate away and do not need to wait
    // const sub = this.client.echo(new EchoRequest({ message: 'text' })).subscribe(res => console.log(res));

    // setTimeout(() => sub.unsubscribe(), 1000); // this closes connection
  }
}
