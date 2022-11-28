import { MesuringPoint } from "../api"

export type MesuringPointsTableProps = {
    points: MesuringPoint[]
    handleDeleteClick: (mesuring_point_id: number) => void
}

const MesuringPointsTable = (props: MesuringPointsTableProps): JSX.Element => {
    // return (<>
    //   {props.points.map(point => <p>{point.id}, {point.body_mass}, {point.created_at}</p>)}
    // </>)
    return (<>
        <table>
            <thead>
                <tr>
                    <td>created</td>
                    <td>mass</td>
                    <td>delete</td>
                </tr>
            </thead>
            <tbody>
                {
                    props.points.map((p) => <tr>
                        <td>{p.created_at}</td>
                        <td>{p.body_mass}</td>
                        <td><button onClick={() => props.handleDeleteClick(p.id)}>delete</button></td>
                    </tr>)
                }
            </tbody>
        </table>
    </>)
}

export default MesuringPointsTable