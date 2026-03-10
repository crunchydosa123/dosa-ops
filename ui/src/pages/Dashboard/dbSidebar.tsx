import React from 'react'
import { Sidebar, SidebarContent, SidebarGroup, SidebarGroupLabel, SidebarHeader, SidebarMenuButton } from '../../components/ui/sidebar'

type Props = {}

const DbSidebar = (props: Props) => {
  return (
    <Sidebar>
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