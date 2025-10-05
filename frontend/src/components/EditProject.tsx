import React, { useState, useEffect } from "react";
import { Dialog, Text, TextField, Flex, Button } from "@radix-ui/themes";
import { models } from "../../wailsjs/go/models";

interface EditProjectProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  project: models.Project;
  onSave: (updatedProject: models.Project) => Promise<void> | void;
}

const EditProject: React.FC<EditProjectProps> = ({ open, onOpenChange, project, onSave }) => {
  const [name, setName] = useState(project.name);
  const [description, setDescription] = useState(project.description);
  const [currency, setCurrency] = useState(project.currency);

  useEffect(() => {
    setName(project.name);
    setDescription(project.description);
    setCurrency(project.currency);
  }, [project]);

  const handleSave = async () => {
    await onSave({ ...project, name, description, currency });
    onOpenChange(false);
  };

  return (
    <Dialog.Root open={open} onOpenChange={onOpenChange}>
      <Dialog.Content maxWidth="480px">
        <Dialog.Title>Edit Project</Dialog.Title>

        <Flex direction="column" gap="3" mt="3">
          <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Name
            </Text>
            <TextField.Root value={name} onChange={(e) => setName(e.target.value)} size="2" />
          </label>

          <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Description
            </Text>
            <TextField.Root value={description} onChange={(e) => setDescription(e.target.value)} size="2" />
          </label>

          <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Currency
            </Text>
            <TextField.Root value={currency} disabled size="2" />
          </label>

          <Flex justify="end" gap="3" mt="4">
            <Button variant="soft" color="gray" onClick={() => onOpenChange(false)}>
              Cancel
            </Button>
            <Button onClick={handleSave}>Save</Button>
          </Flex>
        </Flex>
      </Dialog.Content>
    </Dialog.Root>
  );
};

export default EditProject;

