import axios from 'axios';

declare global {
  interface Window { axios: typeof axios }
}
