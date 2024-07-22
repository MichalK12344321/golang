import { CDG } from "@/api";
import { CollectionStatus } from "@/model";
import { Button } from "primereact/button";
import { useNavigate } from "react-router-dom";

export const getRunElements = (run: CDG.RunDto) : {icon: string, severity: SevType}  => {
  switch (run.status) {
    case CollectionStatus.Created:
      return { icon: "pi pi-hourglass", severity: "help" }
    case CollectionStatus.Started:
      return { icon: "pi pi-spin pi-cog", severity: "info" }
    case CollectionStatus.Success:
      return { icon: "pi pi-check", severity: "success" }
    case CollectionStatus.Terminated:
      return { icon: "pi pi-ban", severity: "warning" }
    case CollectionStatus.Terminating:
      return { icon: "pi pi-times-circle", severity: "warning" }
    case CollectionStatus.Failure:
      return { icon: "pi pi-times", severity: "danger" }
    default:
      return { icon: "", severity: undefined }
  }
}

type Props = {
  run: CDG.RunDto
};

type SevType = 'secondary' | 'success' | 'info' | 'warning' | 'danger' | 'help' | undefined

export function RunIcon({ run }: Props) {
  const navigate = useNavigate();

  const open = (run: CDG.RunDto) => navigate(`/collection/run/${run.runId}`)

  const data = getRunElements(run)
  return (
    <Button key={run.runId} icon={data.icon}  text severity={data.severity} onClick={() => open(run)} />
  );
}
