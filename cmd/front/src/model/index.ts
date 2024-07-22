import { CDG, SCHEDULER } from "@/api";

export interface EventPayload {
  type: EventType;
  data: string // json
}

export interface AppendLineEvent {
  runId: string,
  line: string,
  file: string
}

export interface CollectionCreateEvent extends CDG.CollectionDto {
}

export interface RunCreateEvent extends SCHEDULER.RunScheduleDto {
}

export interface RunUpdateEvent extends CDG.RunDto {
}

export enum CollectionStatus {
  Created = "created",
  Started = "started",
  Failure = "failure",
  Success = "success",
  Terminating = "terminating",
  Terminated = "terminated",
}

export enum EventType {
  CollectionCreateEvent = "CollectionCreateEvent",
  AppendLineEvent = "AppendLineEvent",
  RunCreateEvent = "RunCreateEvent",
  RunUpdateEvent = "RunUpdateEvent"
}

// emitter map
export type EventMap = {
  [EventType.CollectionCreateEvent]: CollectionCreateEvent
  [EventType.AppendLineEvent]: AppendLineEvent
  [EventType.RunCreateEvent]: RunCreateEvent
  [EventType.RunUpdateEvent]: RunUpdateEvent
}

export interface ApiConfig {
  cdgUrl: URL
  schedulerUrl: URL
  webUrl: URL
}