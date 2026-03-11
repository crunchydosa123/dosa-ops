import { SidebarProvider } from "../../components/ui/sidebar"
import DbMain from "./dbMain"
import DbSidebar from "./dbSidebar"

const Dashboard = () => {
  return (
    <div className="h-screen w-screen bg-yellow-200 flex flex-grow">
      <div className="width-1/5">
        <SidebarProvider>
        <DbSidebar />
      </SidebarProvider>
      </div>
      
      <DbMain />
    </div>
  )
}

export default Dashboard