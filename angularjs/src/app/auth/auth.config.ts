import { PassedInitialConfig } from 'angular-auth-oidc-client';

export const authConfig: PassedInitialConfig = {
  config: {
              authority: process.env["APP_AWS_COGNITO_AUTHORITY"],
              redirectUrl: process.env["APP_AWS_COGNITO_REDIRECT_URL"],
              clientId: process.env["APP_AWS_COGNITO_CLIENT_ID"],
              scope: process.env["APP_AWS_COGNITO_SCOPE"],
              responseType: 'code',
              silentRenew: true,
              useRefreshToken: true,
              renewTimeBeforeTokenExpiresInSeconds: 30,
          }
}