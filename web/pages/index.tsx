import { useEffect, useState } from "react";
import { deleteMesuringPoint, getMesuringPoints, MesuringPoint, MesuringPointRequest, postMesuringPoint } from "../api";
import MesuringPointForm from "../components/MesuringPointForm";
import MesuringPointsTable from "../components/MesuringPointsTable";

const USER_ID = 1

export default function Home() {
  const [points, setPoints] = useState<MesuringPoint[]>([])
  const [inputBodyMass, setInputBodyMass] = useState<string>("50")

  const handleClick = async () => {
    const req: MesuringPointRequest = { userId: USER_ID, bodyMass: Number(inputBodyMass) }
    await postMesuringPoint(req)
    setPoints(await getMesuringPoints(USER_ID))
  }

  const handleDeleteClick = async (id: number) => {
    await deleteMesuringPoint(id)
    setPoints(await getMesuringPoints(USER_ID))
  }

  useEffect(
    () => {
      (async () => {
        setPoints(await getMesuringPoints(USER_ID))
      })()
    },
    []
  )

  return (<>
    <MesuringPointForm
      value={inputBodyMass}
      handleClick={handleClick}
      handleChange={(v) => setInputBodyMass(v)}
    />
    <MesuringPointsTable
      points={points}
      handleDeleteClick={handleDeleteClick}
    />
  </>)
}