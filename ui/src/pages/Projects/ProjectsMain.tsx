import React, { useState } from "react"

import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card"
import { Button } from "../../components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogFooter,
} from "../../components/ui/dialog"
import { Input } from "../../components/ui/input"
import { Label } from "../../components/ui/label"

type Project = {
  name: string
  description: string
}

const ProjectsMain = () => {
  const [projects, setProjects] = useState<Project[]>([])
  const [name, setName] = useState("")
  const [description, setDescription] = useState("")

  const addProject = () => {
    if (!name) return

    setProjects([...projects, { name, description }])
    setName("")
    setDescription("")
  }

  return (
    <div className="p-6 flex flex-col gap-6">

      {/* Header */}
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-semibold">Projects</h1>

        <Dialog>
          <DialogTrigger asChild>
            <Button>Add Project</Button>
          </DialogTrigger>

          <DialogContent>
            <DialogHeader>
              <DialogTitle>Create Project</DialogTitle>
            </DialogHeader>

            <div className="flex flex-col gap-4 py-2">
              <div className="flex flex-col gap-1">
                <Label>Project Name</Label>
                <Input
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </div>

              <div className="flex flex-col gap-1">
                <Label>Description</Label>
                <Input
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                />
              </div>
            </div>

            <DialogFooter>
              <Button onClick={addProject}>Create</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      {/* Project Grid */}
      <div className="grid grid-cols-3 gap-4">

        {projects.map((project, i) => (
          <Card key={i} className="hover:shadow-md cursor-pointer">
            <CardHeader>
              <CardTitle>{project.name}</CardTitle>
            </CardHeader>

            <CardContent className="text-sm text-muted-foreground">
              {project.description}
            </CardContent>
          </Card>
        ))}

        {projects.length === 0 && (
          <div className="text-muted-foreground">
            No projects yet. Create one!
          </div>
        )}

      </div>

    </div>
  )
}

export default ProjectsMain