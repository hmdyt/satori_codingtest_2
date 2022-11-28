import axios from "axios"

const HOST = process.env.NEXT_PUBLIC_API_HOST

export type MesuringPoint = {
    id: number
    user_id: number
    body_mass: number
    created_at: string
}

export type MesuringPointRequest = {
    userId: number
    bodyMass: number
}

export type MesuringPointsResponse = {
    mesuring_points: MesuringPoint[]
}

export const getMesuringPoints = async (userId: number): Promise<MesuringPoint[]> => {
    const res = await axios.get<MesuringPointsResponse>(`${HOST}/mesuringpoint/${userId}`)
    return res.data.mesuring_points
}

export const postMesuringPoint = async (req: MesuringPointRequest): Promise<MesuringPoint> => {
    const reqBody = { user_id: req.userId, body_mass: req.bodyMass }
    const res = await axios.post<MesuringPoint>(
        `${HOST}/mesuringpoint`,
        reqBody,
        { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
    )
    return res.data
}

export const deleteMesuringPoint = async (mesuring_point_id: number) => {
    await axios.delete(`${HOST}/mesuringpoint/${mesuring_point_id}`)
}