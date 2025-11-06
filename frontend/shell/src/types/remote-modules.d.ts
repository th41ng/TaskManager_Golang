// ============================================================================
// MODULE FEDERATION TYPE DECLARATIONS
// ============================================================================
// Declare types cho remote modules để TypeScript hiểu props

declare module 'userApp/UserApp' {
  import { ComponentType } from 'react'
  const UserApp: ComponentType<any>
  export default UserApp
}

declare module 'projectApp/ProjectApp' {
  import { ComponentType } from 'react'
  
  interface ProjectAppProps {
    shellData?: {
      user?: any
      theme?: string
      token?: string
    }
    onProjectCreated?: (project: any) => void
    onProjectUpdated?: (project: any) => void
  }
  
  const ProjectApp: ComponentType<ProjectAppProps>
  export default ProjectApp
  
  // Named exports
  export const CreateProjectDialog: ComponentType<any>
  export const EditProjectDialog: ComponentType<any>
}

declare module 'taskApp/TaskApp' {
  import { ComponentType } from 'react'
  
  interface TaskAppProps {
    shellData?: {
      user?: any
      theme?: string
      token?: string
    }
    // optional projectId passed from Shell when user requests to view tasks for a project
    projectId?: number
    onTaskCreated?: (task: any) => void
    onTaskUpdated?: (task: any) => void
  }
  
  const TaskApp: ComponentType<TaskAppProps>
  export default TaskApp
  
  // Named exports
  export const CreateTaskDialog: ComponentType<any>
  export const EditTaskDialog: ComponentType<any>
}
