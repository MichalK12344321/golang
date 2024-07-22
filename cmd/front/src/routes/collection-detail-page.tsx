import { CDG } from "@/api"
import { RunIcon } from "@/components/RunIcon"
import { AppContext } from "@/context/appContext"
import { Button } from "primereact/button"
import { Card } from "primereact/card"
import { ProgressSpinner } from "primereact/progressspinner"
import { useContext, useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"

const CollectionDetail = () => {
  const { client } = useContext(AppContext);
  const { id } = useParams();
  const navigate = useNavigate()

  const [collection, setCollection] = useState<CDG.CollectionDto>()

  useEffect(() => {
    const fetch = async () => {
      try {
        const result = await client.cdgClient.getCollectionId(id ?? "")
        setCollection(result)
      } catch (error) {
        console.error(error)
      }
    }
    fetch()
  }, [])

  const header = () => {
    return <div className="flex flex-column md:flex-row justify-content-between p-2">
      <Button icon="pi pi-chevron-left" label='Collections' onClick={() => navigate("/collection")} />
      <div></div>
    </div>
  }

  const runTemplate = (collection: CDG.CollectionDto) => {
    return collection.runs?.map(x =>
      <div className='col-1' key={x.runId}>
        <RunIcon  run={x} />
      </div>
    )
  }

  if (!collection) {
    return (
      <div className="card flex justify-content-center">
        <ProgressSpinner color='primary' />
      </div>
    );
  }

  return (
    <Card header={header}>
      <span>Id: {collection?.collectionId}</span><br/>
      <span>Type: {collection?.type}</span><br/>
      <span>Script: {collection?.ssh?.script ?? collection?.go?.script}</span>
      <div className="grid">
        {runTemplate(collection!)}
      </div>
    </Card>
  )
}

export default CollectionDetail;