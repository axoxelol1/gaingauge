import { Component, createSignal, ParentComponent, Show } from "solid-js"
import './app.css'
import { A } from "@solidjs/router"

const [isSecretSet, secretSet] = createSignal(localStorage.getItem("secret") != null)
//TODO: add a little button somewhere for clearing secret again in case wrong one is inputted

const App: ParentComponent = props => {
    return (
        < Show when={isSecretSet()} fallback={< SecretPrompt />}>
            <div class='container'>
                <div class='topbar'>
                    <div><A href="/" style="text-decoration: none"><h1>GainGauge 💪</h1></A></div>
                    <div class='navlinks'>
                        <A href="/logging" activeClass="activelink" inactiveClass="inactivelink"> Logging </A>
                        <A href="/exercises" activeClass="activelink" inactiveClass="inactivelink"> Exercises </A>
                    </div>
                </div>
                <div>
                    {props.children}
                </div>
            </div>
        </Show >
    )
}


const SecretPrompt: Component = () => {
    const [secret, setSecret] = createSignal("")
    const saveSecret = () => {
        localStorage.setItem("secret", secret())
        secretSet(true)
    }
    return <>
        <label> Secret: <input type="text" name="secret" value={secret()} onChange={(e) => setSecret(e.currentTarget.value)} /> </label>
        <button onClick={saveSecret}> Submit </button>
    </>
}

export default App
