import { httpService } from '../src/services/httpService'
var bodyParser = require('body-parser')


function App() {
  var a = httpService.get('ungroup','rrrrrrrrrrrrrr');
  a.then(val => console.log(val))
  console.log(a);

  return (
    <div className="App"></div>
  );
}

export default App;
