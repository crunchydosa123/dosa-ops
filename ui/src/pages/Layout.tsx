import React from "react";
import { SidebarProvider } from "../components/ui/sidebar";
import DbSidebar from "./Dashboard/dbSidebar";

type Props = {
  children: React.ReactNode;
};

const Layout = ({ children }: Props) => {
  return (
    <SidebarProvider>
      <div className="h-screen w-full flex border border-black bg-yellow-200">
        <DbSidebar />

        <div className="flex-1 h-full">
          {children}
        </div>
      </div>
    </SidebarProvider>
  );
};

export default Layout;