import { Component, createResource, Match, Show, Switch } from "solid-js";

const fetchUsers = async () =>
    (await fetch('/api/users')).json();


const Dashboard: Component = (props) => {
    const [users] = createResource(fetchUsers)
    return (
        <>
            <h1> Dashboard </h1>
            <div>
                <Show when={users.loading}>
                    <p>Loading...</p>
                </Show>
                <Switch>
                    <Match when={users.error}>
                        <span>Error {users.loading}</span>
                    </Match>
                    <Match when={users()}>
                        <div>{JSON.stringify(users())}</div>
                    </Match>
                </Switch>
            </div>
        </>
    )
}

export default Dashboard
