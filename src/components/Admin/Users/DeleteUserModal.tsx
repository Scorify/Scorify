import { ConfirmModal } from "../..";

type props = {
  user: string;
  open: boolean;
  setOpen: (open: boolean) => void;
  handleDelete: () => void;
};

export default function DeleteUserModal({
  user,
  open,
  setOpen,
  handleDelete,
}: props) {
  return (
    <ConfirmModal
      title='Delete User'
      subtitle={
        <>
          To confirm deletion of user, type the name (<b>{user}</b>) of the user
          below.
        </>
      }
      buttonText='Delete User'
      value={user}
      open={open}
      setOpen={setOpen}
      onConfirm={handleDelete}
      label='User Name'
    />
  );
}
