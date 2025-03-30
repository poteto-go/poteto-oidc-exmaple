import axios from "redaxios";

export class OidcClient {
  private _clientId: string;
  private _authEndpoint: string;
  private _redirectUrl: string;

  constructor(props: {
    clientId: string;
    authEndpoint: string;
    redirectUrl: string;
  }) {
    this._clientId = props.clientId;
    this._authEndpoint = props.authEndpoint;
    this._redirectUrl = props.redirectUrl;
  }

  public getAuthUrl(props: {
    respType: string;
    scopes: string[];
    state: string;
  }): string {
    const { respType, scopes, state } = props;

    return `${this._authEndpoint}?client_id=${
      this._clientId
    }&response_type=${respType}&scope=${scopes.join("%20")}&redirect_uri=${
      this._redirectUrl
    }&state=${state}`;
  }
}
