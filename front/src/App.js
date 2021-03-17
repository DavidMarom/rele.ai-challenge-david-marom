import { httpService } from '../src/services/httpService'


function App() {

  let data = `
type: DataFileWrite
payload:
    a: "foo"
    b: "/tmp"
---
type: DataFileTimeout
payload:
    c: 60
---
type: DataFileTimeout
payload:
    d: 880
    e: 4888
    f: 345888
`;
  var a = httpService.get('ungroup', data);
  a.then(val => console.log(val))

//   let data2 = `
// type: DataFile
// payload:
// content: "foo"
// destination: "/tmp"
// timeout: 60
// `;
// var b = httpService.get('group', data2);
//   b.then(val => console.log(val))


  return (
    <div className="App"></div>
  );
}

export default App;
