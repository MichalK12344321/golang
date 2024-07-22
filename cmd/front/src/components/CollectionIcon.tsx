import { CDG } from "@/api";
import { Button } from "primereact/button";
import { useNavigate } from "react-router-dom";


type Props = {
  value: CDG.CollectionDto
};

export function CollectionIcon({ value: value }: Props) {
  const navigate = useNavigate();

  const open = (v: CDG.CollectionDto) => navigate(`/collection/${v.collectionId}`)

  return (
    <Button icon="pi pi-align-left" key={value.collectionId} text onClick={() => open(value)} />
  );
}
