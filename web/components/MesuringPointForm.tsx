export type MesuringPointFormProps = {
    value: string
    handleChange: (v: string) => void
    handleClick: () => void
}

const MesuringPointForm = (props: MesuringPointFormProps): JSX.Element => {
    return (
        <>
            <input type={"number"} value={props.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => props.handleChange(e.target.value)}></input>
            <button onClick={props.handleClick}>submit</button>
        </>
    )
}

export default MesuringPointForm