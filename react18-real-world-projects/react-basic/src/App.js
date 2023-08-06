import { useState } from "react";

function App() {
  const [count, setCount] = useState(0);

  function addCount() {
    setCount(count + 1);
  }

  function downCount() {
    setCount(count - 1);
  }

  function resetCount() {
    setCount(0);
  }

  return (
    <>
      <h1>{count}</h1>
      <button onClick={addCount}>add</button>
      <button onClick={downCount}>down</button>
      <button onClick={resetCount}>reset</button>
    </>
  );
}

export default App;
