import { bootstrapApplication } from '@angular/platform-browser';
import { AppConfigModule } from './app/app.config';
import { AppComponent } from './app/app.component';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

// bootstrapApplication(AppComponent, appConfig).catch(err => console.error(err));

platformBrowserDynamic()
  .bootstrapModule(AppConfigModule)
  .catch(err => console.error(err));
