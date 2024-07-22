import { Button } from "primereact/button";

type Props = {
  label?: string,
  copyValue: string
};

export function ClipboardButton({ label, copyValue }: Props) {
  const save = () => {
    navigator.clipboard.writeText(copyValue);
  }
  return (
    <Button  label={label} icon="pi pi-copy" iconPos="right" text onClick={save} />
  );
}
