import { useContext, useEffect } from "react";
import { AppContext } from "./appContext";
import { EventPayload, EventMap } from "@/model";
import useWebSocket from "react-use-websocket";
import { trimEnd } from 'lodash'

interface Props {
  children: React.ReactNode;
}

export const AppContextProvider: React.FunctionComponent<Props> = (props: Props) => {
  const { client, eventBus, config } = useContext(AppContext);

  const socketUrl = `ws://${config.webUrl.host}${trimEnd(config.webUrl.pathname, "/")}/event`
  const { lastMessage } = useWebSocket(socketUrl);

  useEffect(() => {
    if (lastMessage !== null) {
      const msg = lastMessage as MessageEvent<string>
      const brokerEvent: EventPayload = JSON.parse(msg.data)
      let eventData: any = {}
      try {
        eventData = JSON.parse(brokerEvent.data)
      } catch (error) {
        console.error(error)
      }
      eventBus.emit(brokerEvent.type as keyof EventMap, eventData as any)
    }
  }, [lastMessage]);

  return <AppContext.Provider value={{ client, eventBus: eventBus, config }}>
    {props.children}
  </AppContext.Provider>
}