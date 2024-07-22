const envVariables = import.meta.env;

interface Env {
  SCHEDULER_URL: string
  CDG_URL: string
  WEB_URL: string
}

export const envs: Env = {
  SCHEDULER_URL: envVariables.VITE_SCHEDULER_URL,
  CDG_URL: envVariables.VITE_CDG_URL,
  WEB_URL: envVariables.VITE_WEB_URL,
};
