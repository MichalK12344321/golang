import { axiosInstance, CDG, SCHEDULER } from "@/api";
import React from "react";
import mitt, { Emitter } from 'mitt'
import { ApiConfig, EventMap } from "@/model";
import { envs } from "@/envs";
import { trimEnd } from "lodash";

export interface AppState {
  config: ApiConfig,
  client: {
    cdgClient: CDG.IClient,
    schedulerClient: SCHEDULER.IClient
  }
  eventBus: Emitter<EventMap>
}

const config: ApiConfig = {
  schedulerUrl: new URL(envs.SCHEDULER_URL, window.location.toString()),
  cdgUrl: new URL(envs.CDG_URL, window.location.toString()),
  webUrl: new URL(envs.WEB_URL, window.location.toString()),
}

const defaultState: AppState = {
  config: config,
  client: {
    schedulerClient: new SCHEDULER.Client(trimEnd(config.schedulerUrl.toString(), "/"), axiosInstance),
    cdgClient: new CDG.Client(trimEnd(config.cdgUrl.toString(), "/"), axiosInstance),
  },
  eventBus: mitt<EventMap>()
}

export const AppContext = React.createContext<AppState>(defaultState);