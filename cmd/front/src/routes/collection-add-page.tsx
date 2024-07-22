import { useContext, useState } from 'react';
import { AppContext } from '@/context/appContext';
import { SCHEDULER } from '@/api';
import { InputText } from 'primereact/inputtext';
import { SelectButton } from 'primereact/selectbutton';
import { useNavigate } from 'react-router-dom';
import { Card } from 'primereact/card';
import { Button } from 'primereact/button';
import { ToggleButton } from 'primereact/togglebutton';
import { FloatLabel } from "primereact/floatlabel";
import { InputTextarea } from 'primereact/inputtextarea';

const CollectionAdd = () => {
  const navigate = useNavigate();
  const { client } = useContext(AppContext);

  const [typeOptions] = useState([
    { label: 'SSH', value: 'ssh' },
    { label: 'Go', value: 'go' },
  ]);

  const [type, setType] = useState<"ssh" | "go">("ssh")
  const [hasCron, setHasCron] = useState(false)

  const [scheduleOptions, setScheduleOptions] = useState<{ cron: string, repeat: boolean }>({ cron: "* * * * *", repeat: false })

  const newTarget = () => {
    return {
      host: "ne",
      port: 22,
      user: "root",
      password: "root",
    } as SCHEDULER.TargetDto
  }

  const [sshOptions, setSshOptions] = useState<SCHEDULER.ScheduleSSHCollectionDto>({
    script: "/scripts/infinite.sh",
    targets: [newTarget()],
    timeout: "1m"
  })

  const [goOptions, setGoOptions] = useState<SCHEDULER.ScheduleGoCollectionDto>({
    script:
      `package main

import "fmt"

func main(){
  fmt.Println("out")
}
`,
    timeout: "5s"
  })

  const schedule = async () => {
    try {
      let runs: SCHEDULER.RunScheduleDto[] = []

      if (type == "ssh") {
        sshOptions.cron = hasCron ? scheduleOptions.cron : ""
        sshOptions.repeat = scheduleOptions.repeat
        const result = await client.schedulerClient.postScheduleSsh(sshOptions)
        runs = result.runs
      }

      if (type == "go") {
        goOptions.cron = hasCron ? scheduleOptions.cron : ""
        goOptions.repeat = scheduleOptions.repeat
        const result = await client.schedulerClient.postScheduleGo(goOptions)
        runs = result.runs
      }

      if (runs.length == 1) {
        navigate(`/collection/run/${runs[0].runId}`)
      }
      else {
        navigate(`/collection`)
      }
    } catch (error) {
      console.error(error)
    }
  }

  const header = () => {
    return <div className="flex flex-column md:flex-row justify-content-between p-2">
      <Button icon="pi pi-chevron-left" label='Collections' onClick={() => navigate("/collection")} />
      <div></div>
      <Button icon="pi pi-hourglass" label='Schedule' onClick={schedule} />
    </div>
  }

  const sshTargetTemplate = (target: SCHEDULER.TargetDto, index: number) => {
    return <div key={index} className="grid py-2">
      <div className='col-1'>
        <Button disabled={index == 0} icon="pi pi-times" text onClick={() => {
          const arr = sshOptions.targets
          sshOptions.targets.splice(index, 1)
          setSshOptions({ ...sshOptions, targets: arr })
        }} />
      </div>
      <div className='col-3'>
        <FloatLabel>
          <InputText id={`ssh-target-${index}-host`} placeholder="Host" value={target.host} style={{ width: "100%" }} onChange={(e) => {
            const result = sshOptions.targets.slice()
            result[index].host = e.target.value
            setSshOptions({ ...sshOptions, targets: result })
          }} />
          <label htmlFor={`ssh-target-${index}-host`}>Host</label>
        </FloatLabel>
      </div>
      <div className='col-2'>
        <FloatLabel>
          <InputText id={`ssh-target-${index}-port`} placeholder="Port" value={target.port + ""} style={{ width: "100%" }} onChange={(e) => {
            const result = sshOptions.targets.slice()
            result[index].port = +e.target.value
            setSshOptions({ ...sshOptions, targets: result })
          }} />
          <label htmlFor={`ssh-target-${index}-port`}>Port</label>
        </FloatLabel>
      </div>
      <div className='col-3'>
        <FloatLabel>
          <InputText id={`ssh-target-${index}-user`} placeholder="User" value={target.user} style={{ width: "100%" }} onChange={(e) => {
            const result = sshOptions.targets.slice()
            result[index].user = e.target.value
            setSshOptions({ ...sshOptions, targets: result })
          }} />
          <label htmlFor={`ssh-target-${index}-user`}>User</label>
        </FloatLabel>
      </div>
      <div className='col-3'>
        <FloatLabel>
          <InputText id={`ssh-target-${index}-pass`} placeholder="Password" type='password' value={target.password} style={{ width: "100%" }} onChange={(e) => {
            const result = sshOptions.targets.slice()
            result[index].password = e.target.value
            setSshOptions({ ...sshOptions, targets: result })
          }} />
          <label htmlFor={`ssh-target-${index}-pass`}>Password</label>
        </FloatLabel>
      </div>
    </div>
  }

  const scheduleTemplate = () => {
    return (
      <FloatLabel>
        <InputText id="schedule-cron" placeholder="Cron" value={scheduleOptions.cron} style={{ width: "100%" }} onChange={e => setScheduleOptions({ ...scheduleOptions, cron: e.target.value })} />
        <label htmlFor="schedule-cron">Cron</label>
      </FloatLabel>
    )
  }

  const sshTemplate = () => {
    return (
      <div className="grid">
        <div className='col-12 py-4'>
          <FloatLabel>
            <InputText id="ssh-timeout" placeholder="Script" value={sshOptions.timeout} style={{ width: "100%" }} onChange={e => setSshOptions({ ...sshOptions, timeout: e.target.value })} />
            <label htmlFor="ssh-timeout">Timeout</label>
          </FloatLabel>
        </div>
        <div className='col-12'>
          <FloatLabel>
            <InputText id="ssh-script" placeholder="Script" value={sshOptions.script} style={{ width: "100%" }} onChange={e => setSshOptions({ ...sshOptions, script: e.target.value })} />
            <label htmlFor="ssh-script">Script</label>
          </FloatLabel>
        </div>

        <div className='col-12 mt-4'>
          <FloatLabel>
            <label htmlFor="ssh-targets">Targets</label>
            <div id="ssh-targets">
              {sshOptions.targets.map(sshTargetTemplate)}
            </div>
          </FloatLabel>
        </div>

        <div className='col-1'>
          <Button icon="pi pi-plus" text onClick={() => setSshOptions({ ...sshOptions, targets: [...sshOptions.targets, newTarget()] })} />
        </div>
      </div>
    )
  }

  const goTemplate = () => {
    return (
      <div className="grid">
        <div className='col-12 py-4'>
          <FloatLabel>
            <InputText id="go-timeout" value={goOptions.timeout} style={{ width: "100%" }} />
            <label htmlFor="go-timeout">Timeout</label>
          </FloatLabel>
        </div>
        <div className='col-12'>
          <FloatLabel>
            <InputTextarea style={{ width: "100%", height: "300px" }} id="go-script" value={goOptions.script} onChange={(e) => setGoOptions({ ...goOptions, script: e.target.value })} />
            <label htmlFor="go-script">Script</label>
          </FloatLabel>
        </div>
      </div>
    )
  }

  return (
    <Card header={header}>
      <div className="grid">
        <div className='col-2'>
          <SelectButton value={type} onChange={(e) => setType(e.value)} options={typeOptions} />
        </div>
        <div className='col-6'></div>
        <div className='col-2'>
          <ToggleButton onLabel="Scheduled" offLabel='Immediate' checked={hasCron} onChange={e => setHasCron(e.target.value)} />
        </div>
        <div className='col-2'>
          <ToggleButton disabled={!hasCron} onLabel="Repeat" offLabel='One time' checked={scheduleOptions.repeat} onChange={() => setScheduleOptions({ ...scheduleOptions, repeat: !scheduleOptions.repeat })} />
        </div>
        {hasCron ? <div className='col-12 mt-4'>
          {scheduleTemplate()}
        </div> : undefined}
        <div className='col-12'>
          {type == "ssh" ? sshTemplate() : goTemplate()}
        </div>
      </div>
    </Card>
  )
}

export default CollectionAdd;