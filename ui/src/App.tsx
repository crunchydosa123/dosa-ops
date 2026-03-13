import { BrowserRouter  as Router, Routes, Route} from "react-router";
import Home from './pages/Home';
import Dashboard from "./pages/Dashboard/Dashboard";
import Projects from "./pages/Projects/Projects";
import SingleProject from "./pages/SingleProject/SingleProject";
import Login from "./pages/Auth/Login";
import Signup from "./pages/Auth/Signup";


const App = () => {
  return (
    <Router>
   <Routes>
    <Route path='/' element={<Home/>}/>

    <Route path='/dashboard' element={<Dashboard />}/>
    <Route path='/projects' element={<Projects />}/>
    <Route path='/project/:id' element={<SingleProject />}/>

     <Route path='/auth/login' element={<Login />}/>
     <Route path='/auth/signup' element={<Signup />}/>
     
   </Routes>
   </Router>
  )
}

export default App