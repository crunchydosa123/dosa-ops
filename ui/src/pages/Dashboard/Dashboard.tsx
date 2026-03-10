import { SidebarProvider } from "../../components/ui/sidebar"
import DbSidebar from "./dbSidebar"

const Dashboard = () => {
  return (
    <div className="h-screen w-screen bg-yellow-200">
      <SidebarProvider>
        <DbSidebar />
      </SidebarProvider>
    </div>
  )
}

export default Dashboard