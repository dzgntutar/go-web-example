import "./App.css";
import { BrowserRouter as Router, Route, Switch ,Link} from "react-router-dom";
import Home from "./pages/Home";
import Posts from "./components/Posts";
import Categories from "./components/Categories"

function App() {
  return (
    <div className="container">
      <Router>
      <div className="row mt-2">
        <div className="col-md-4">
          <ul className="list-group">
            <li className="list-group-item "><Link to="/"> Home </Link></li>
            <li className="list-group-item "><Link to="/post"> Post </Link></li>
            <li className="list-group-item"><Link to="/category">Category</Link></li>
          </ul>

        </div>
        <div className="col-md-8">
          
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/post" component={Posts} />
              <Route exact path="/category" component={Categories} />

            </Switch>
          
        </div>
      </div>
      </Router>
    </div>
  );
}

export default App;
