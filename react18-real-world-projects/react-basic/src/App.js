import { useState } from "react";

function App() {
  const [students, setStudent] = useState([
    { id: "1", name: "Ant" },
    { id: "2", name: "Bird" },
    { id: "3", name: "Cat" },
    { id: "4", name: "Dog" },
  ]);

  const [show, setShow] = useState(true);

  function deleteStudent(id) {
    setStudent(students.filter((stud) => stud.id !== id));
  }

  return (
    <>
      <h1>มีนักเรียนทั้งหมด : {students.length}</h1>

      <h1>{show ? "SHOW" : "OFF"}</h1>
      <button onClick={() => setShow(!show)}>Toggle</button>

      <ul>
        {show &&
          students.map((stud) => (
            <li key={stud.id}>
              <p>
                {stud.id} : {stud.name}
              </p>
              <button onClick={() => deleteStudent(stud.id)}>delete</button>
            </li>
          ))}
      </ul>
    </>
  );
}

export default App;
