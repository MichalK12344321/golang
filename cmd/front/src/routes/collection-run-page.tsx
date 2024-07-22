import { useContext, useEffect, useState } from 'react';
import { AppContext } from '@/context/appContext';
import { CDG } from '@/api';
import { useNavigate, useParams } from 'react-router-dom';
import { Card } from 'primereact/card';
import { Button } from 'primereact/button';
import { ButtonGroup } from 'primereact/buttongroup';
import useWebSocket from 'react-use-websocket';
import { ScrollPanel } from 'primereact/scrollpanel';
import { AppendLineEvent, CollectionStatus, EventPayload, EventType, RunUpdateEvent } from '@/model';
import { useMountEffect } from 'primereact/hooks';
import { RunIcon } from '@/components/RunIcon';
import { ClipboardButton } from '@/components/ClipboardButton';
import { trimEnd } from 'lodash';
import { ProgressSpinner } from 'primereact/progressspinner';


const CollectionRun = () => {
  const navigate = useNavigate();
  const { client, eventBus, config } = useContext(AppContext);

  const { id } = useParams();
  const [run, setRun] = useState<CDG.RunDto>()

  const handleRunUpdate = (event: RunUpdateEvent) => {
    if (event.runId == id) {
      setRun(run => {
        if (run) {
          return Object.assign({}, event)
        }
      })
    }
  }

  useMountEffect(() => {
    eventBus.on<EventType.RunUpdateEvent>(EventType.RunUpdateEvent, handleRunUpdate)
  })

  // useUnmountEffect(() => {
  //   eventBus.off<EventType.RunUpdateEvent>(EventType.RunUpdateEvent, handleRunUpdate)
  // })

  useEffect(() => {
    const fetch = async () => {
      try {
        const result = await client.cdgClient.getCollectionRunRunId(id ?? "")
        if (!run?.collectionId) { // update event may win the race
          setRun(result)
        }
      } catch (error) {
        console.error(error)
      }
    }

    fetch()

  }, [])

  const socketUrl = `ws://${config.webUrl.host}${trimEnd(config.webUrl.pathname, "/")}/run/${id}/preview `
  const [messageHistory, setMessageHistory] = useState<AppendLineEvent[]>([]);
  const { lastMessage } = useWebSocket(socketUrl);

  useEffect(() => {
    if (lastMessage !== null) {
      const msg = lastMessage as MessageEvent<string>
      const p: EventPayload = JSON.parse(msg.data)
      setMessageHistory((prev) => prev.concat(JSON.parse(p.data)));
    }
  }, [lastMessage]);

  const terminate = async () => {
    await client.schedulerClient.postTerminate({ runId: run?.runId ?? "" })
  }

  const getArchive = () => {
    window.open(`${new URL(config.cdgUrl)}/collection/run/${id}/archive`, "_blank")
  }

  const header = () => {
    return <div className="flex flex-column md:flex-row justify-content-between p-2">
      <Button icon="pi pi-chevron-left" label='Collections' onClick={() => navigate("/collection")} />
      <ButtonGroup>
        <RunIcon run={run!} />
        <ClipboardButton label={id ?? ""} copyValue={id ?? ""} />
      </ButtonGroup>
      <div></div>
      <Button disabled={run?.status == CollectionStatus.Started} icon="pi pi-download" label='Archive' onClick={getArchive} />
      <Button disabled={run?.status != CollectionStatus.Started} icon="pi pi-ban" label='Terminate' onClick={terminate} />
    </div>
  }

  if (!run?.runId) {
    return (
      <div className="card flex justify-content-center">
        <ProgressSpinner color='primary' />
      </div>
    );
  }

  return (
    <Card header={header}>
      {messageHistory.length > 0 ?
        <ScrollPanel style={{ width: '100%', height: '50vh' }}>
          {messageHistory.map((message, idx) => (
            <div key={idx}>
              <span className="p-terminal-prompt text-gray-400 mr-2" key={idx}>{message.line}</span>
              <br />
            </div>
          ))}
          {run?.status == CollectionStatus.Started ? <i className="pi pi-spin pi-spinner"></i> : <div></div>}
        </ScrollPanel>
        : run?.status == CollectionStatus.Success ? <span>preview unavailable, please download archive</span> : undefined}
      {run.error ? <span className="p-terminal-prompt text-red-400 mr-2">error: {run.error}</span> : undefined}
    </Card>
  )
}

export default CollectionRun;