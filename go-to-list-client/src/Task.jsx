import { useState } from "react";

export default function Task({ task, i, handleDelete, handleEdit }) {

    const [isEdit, setIsEdit] = useState(false);
    const [editedTaskName, setEditedTaskName] = useState(task);

    return (
        <div className="p-3 border-bottom d-flex align-items-center">
            {
                isEdit ?
                    <form onSubmit={e =>handleEdit(e, i, editedTaskName, setIsEdit)}>
                        <input value={editedTaskName} onChange={e=>setEditedTaskName(e.target.value)} type="text" className="form-control me-5" />
                    </form>
                    :
                    <h6 className="f-600 mb-0">{task}</h6>
            }
            <button onClick={() => setIsEdit(true)} className="btn btn-sm btn-outline-secondary ms-auto">Edit</button>
            <button onClick={() => handleDelete(i)} className="btn btn-sm btn-outline-danger ms-2">Delete</button>
        </div>
    );
}