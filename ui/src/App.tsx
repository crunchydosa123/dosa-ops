import { BrowserRouter  as Router, Routes, Route} from "react-router";
import Home from './pages/Home';
import Dashboard from "./pages/Dashboard/Dashboard";


const App = () => {
  return (
    <Router>
   <Routes>
    <Route path='/' element={<Home/>}/>

    <Route path='/dashboard' element={<Dashboard />}/>
     
   </Routes>
   </Router>
  )
}

export default App