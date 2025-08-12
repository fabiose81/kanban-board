import { ApplicationConfig, provideBrowserGlobalErrorListeners, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { OidcSecurityService } from 'angular-auth-oidc-client'
import { firstValueFrom } from 'rxjs';

import { routes } from './app.routes';
import { authConfig } from './auth/auth.config';
import { provideAuth } from 'angular-auth-oidc-client';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideHttpClient(),
    provideAuth(authConfig),
    {
      provide: 'app-init',
      multi: true,
      useFactory: (oidc: OidcSecurityService) => {
        return async () => {
          await firstValueFrom(oidc.checkAuth());
        };
      },
    }
  ]
};
