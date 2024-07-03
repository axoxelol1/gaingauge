import { Component, createResource, Match, Switch } from "solid-js";
import { User } from "../utils/types";

const fetchUsers: () => Promise<User[]> = async () => {
    // TODO: api error handling
    const res = await fetch('/api/users', { headers: { "Authorization": localStorage.getItem("secret") } });
    if (!res.ok) {
        throw new Error(res.statusText)
    }
    return res.json() as Promise<User[]>
}


const Dashboard: Component = () => {
    const [users] = createResource(fetchUsers)
    console.log(users)
    return (
        <>
            <Switch>
                <Match when={users.error}>
                    {users.error}
                </Match>
                <Match when={users.loading}>
                    Loading...
                </Match>
                <Match when={users()}>
                    <div>{JSON.stringify(users())}</div>
                </Match>
            </Switch>
        </>
    )
}

export default Dashboard
