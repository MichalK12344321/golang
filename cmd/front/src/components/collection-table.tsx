import { useContext, useEffect, useState } from 'react';
import { AppContext } from '@/context/appContext';
import { CDG } from '@/api';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Button } from 'primereact/button';
import { useNavigate } from 'react-router-dom';
import { getRunElements, RunIcon } from '@/components/RunIcon';
import { useMountEffect } from 'primereact/hooks';
import { CollectionCreateEvent, CollectionStatus, EventType, RunUpdateEvent } from '@/model';
import { CollectionIcon } from './CollectionIcon';
import { MultiSelect } from 'primereact/multiselect';

const CollectionTable = () => {
  const fetchCount = 50;
  const navigate = useNavigate();

  const { client, eventBus } = useContext(AppContext);

  const [collections, setCollections] = useState<CDG.CollectionDto[]>([])
  const [statuses, setStatuses] = useState<CollectionStatus[]>([])

  const handleCreateCollection = async (event: CollectionCreateEvent) => {
    try {
      const newCollection = await client.cdgClient.getCollectionId(event.collectionId!)
      setCollections(p => {
        if (!p.find(x => x.collectionId == newCollection.collectionId)) {
          return [newCollection, ...p]
        }
        return p
      })
    } catch (error) {
      console.error(error)
    }
  }

  const handleRunUpdate = (event: RunUpdateEvent) => {
    setCollections(collections => {
      const result = collections.map(x => Object.assign({}, x))
      const collection = result.find(x => x.collectionId == event.collectionId)
      if (collection) {
        const run = collection.runs?.find(x => x.runId == event.runId)
        if (run) {
          run.status = event.status
          run.error = event.error
        }
        else {
          collection.runs = [event, ...collection.runs ?? []]
        }
      }
      return result
    })
  }

  useMountEffect(() => {
    eventBus.on<EventType.CollectionCreateEvent>(EventType.CollectionCreateEvent, handleCreateCollection)
    eventBus.on<EventType.RunUpdateEvent>(EventType.RunUpdateEvent, handleRunUpdate)
  })

  // useUnmountEffect(() =>{
  //   eventBus.off<EventType.CollectionCreateEvent>(EventType.CollectionCreateEvent, handleCreateCollection)
  //   eventBus.off<EventType.RunUpdateEvent>(EventType.RunUpdateEvent, handleRunUpdate)
  // })

  useEffect(() => {
    const fetch = async () => {
      try {
        const st = statuses.length > 0 ? statuses.join(",") : undefined;
        const collections = await client.cdgClient.getCollection(fetchCount, undefined, st)
        setCollections(collections)
      } catch (error) {
        console.error(error)
      }
    }
    fetch()
  }, [statuses])

  const idTemplate = (collection: CDG.CollectionDto) => {
    return <CollectionIcon key={collection.collectionId} value={collection} />
  }

  const runTemplate = (collection: CDG.CollectionDto) => {
    return collection.runs?.map(x => <RunIcon key={x.runId} run={x} />)
  }

  const header = (
    <div className="flex flex-wrap align-items-center justify-content-between pa-2">
      <div></div>
      <Button icon="pi pi-plus" label='Add' onClick={() => navigate("add")} />
    </div>
  );

  const statusFilterItemTemplate = (props: any) => {
    const option = props as CollectionStatus
    const d = getRunElements({status: option})
    return <span>{option} <i className={d.icon}></i></span>
  }

  const statusFilterTemplate = () => {
    return (
      <MultiSelect value={statuses} options={Object.values(CollectionStatus)}
        placeholder="Filter by status" className="p-column-filter" itemTemplate={statusFilterItemTemplate} showClear onChange={(e) => {
          setStatuses(e.target.value)
        }} />
    );
  };

  return (
    <div className="card">
      <DataTable value={collections} filters={{}} filterDisplay="row" header={header} paginator rows={5} rowsPerPageOptions={[5, 10, 25, 50]} tableStyle={{ minWidth: '50rem' }} >
        <Column body={idTemplate} field="collectionId" header="ID"></Column>
        <Column field="type" header="Type"></Column>
        <Column body={runTemplate} header="Runs" filter filterElement={statusFilterTemplate} showFilterMenu={false} filterMenuStyle={{ width: '14rem' }} style={{ minWidth: '12rem' }}></Column>
      </DataTable>
    </div>
  );
}

export default CollectionTable;