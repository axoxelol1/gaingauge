import { ParentComponent } from "solid-js"

const App: ParentComponent = props => (
    <>
        <h1>Site Title</h1>
        {props.children}
    </>
)

export default App
