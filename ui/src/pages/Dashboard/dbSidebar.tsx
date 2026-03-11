import { Sidebar, SidebarContent, SidebarGroup, SidebarGroupLabel, SidebarHeader, SidebarMenuButton } from '../../components/ui/sidebar'

const DbSidebar = () => {
  return (
    <Sidebar className=''>
      <SidebarHeader>
        Dosa-Ops
      </SidebarHeader>
      <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>
              Main Menu
            </SidebarGroupLabel>
            <SidebarMenuButton>Projects</SidebarMenuButton>
            <SidebarMenuButton>Applications</SidebarMenuButton>
          </SidebarGroup>
        </SidebarContent>

    </Sidebar>

  )
}

export default DbSidebar