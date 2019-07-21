package com.github.zxh.classpy.gui;

import com.github.zxh.classpy.gui.jar.JarTreeView;
import com.github.zxh.classpy.gui.parsed.ParsedViewerPane;
import com.github.zxh.classpy.gui.support.*;
import com.github.zxh.classpy.helper.UrlHelper;
import javafx.application.Application;
import javafx.beans.value.ObservableValue;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.input.Dragboard;
import javafx.scene.input.TransferMode;
import javafx.scene.layout.BorderPane;
import javafx.scene.text.Text;
import javafx.stage.Stage;

import java.io.File;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.Optional;

/**
 * Main class.
 */
public class ClasspyApp extends Application {

    private static final String TITLE = "Classpy";


    private Stage stage;
    private BorderPane root;
    private MyMenuBar menuBar;

    @Override
    public void start(Stage stage) {
        this.stage = stage;

        root = new BorderPane();
        root.setTop(createMenuBar());
        root.setCenter(createTabPane());

        Scene scene = new Scene(root, 960, 540);
        //scene.getStylesheets().add("classpy.css");
        enableDragAndDrop(scene);

        stage.setScene(scene);
        stage.setTitle(TITLE);
        stage.getIcons().add(ImageHelper.loadImage("/spy16.png"));
        stage.getIcons().add(ImageHelper.loadImage("/spy32.png"));
        stage.show();
    }

    private TabPane createTabPane() {
        TabPane tp = new TabPane();
        tp.getSelectionModel().selectedItemProperty().addListener(
                (ObservableValue<? extends Tab> observable, Tab oldTab, Tab newTab) -> {
                    if (newTab != null) {
                        URL url = (URL) newTab.getUserData();
                        stage.setTitle(TITLE + " - " + url);
                    }
        });
        return tp;
    }

    private Tab createTab(URL url) {
        Tab tab = new Tab();
        tab.setText(UrlHelper.getFileName(url));
        tab.setUserData(url);
        tab.setContent(new BorderPane(new ProgressBar()));
        ((TabPane) root.getCenter()).getTabs().add(tab);
        return tab;
    }

    private MenuBar createMenuBar() {
        menuBar = new MyMenuBar();

        menuBar.setOnOpenFile(this::onOpenFile);
        menuBar.setOnNewWindow(this::openNewWindow);
        //menuBar.setUseSystemMenuBar(true);

        return menuBar;
    }

    // http://www.java2s.com/Code/Java/JavaFX/DraganddropfiletoScene.htm
    private void enableDragAndDrop(Scene scene) {
        scene.setOnDragOver(event -> {
            Dragboard db = event.getDragboard();
            if (db.hasFiles()) {
                event.acceptTransferModes(TransferMode.COPY);
            } else {
                event.consume();
            }
        });

        // Dropping over surface
        scene.setOnDragDropped(event -> {
            Dragboard db = event.getDragboard();
            boolean success = false;
            if (db.hasFiles()) {
                success = true;
                for (File file : db.getFiles()) {
                    //System.out.println(file.getAbsolutePath());
                    openFile(file);
                }
            }
            event.setDropCompleted(success);
            event.consume();
        });
    }

    private void openNewWindow() {
        ClasspyApp newApp = new ClasspyApp();
        // is this correct?
        newApp.start(new Stage());
    }

    private void onOpenFile(FileType ft, URL url) {
        if (url == null) {
            if (ft == FileType.BITCOIN_BLOCK) {
                showBitcoinBlockDialog();
            } else if (ft == FileType.BITCOIN_TX) {
                showBitcoinTxDialog();
            } else {
                showFileChooser(ft);
            }
        } else {
            openFile(url);
        }
    }

    private void showBitcoinBlockDialog() {
        String apiUrl = "https://blockchain.info/rawblock/<hash>?format=hex";
        String genesisBlockHash = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f";

        TextInputDialog dialog = new TextInputDialog(genesisBlockHash);
        dialog.setTitle("Block Hash Input Dialog");
        dialog.setHeaderText("API: " + apiUrl);
        dialog.setContentText("hash: ");
        dialog.setResizable(true);

        // Traditional way to get the response value.
        Optional<String> result = dialog.showAndWait();
        if (result.isPresent()){
            try {
                openFile(new URL(apiUrl.replace("<hash>", result.get())));
            } catch (MalformedURLException ignored) {}
        }
    }

    private void showBitcoinTxDialog() {
        String apiUrl = "https://blockchain.info/rawtx/<hash>?format=hex";

        TextInputDialog dialog = new TextInputDialog();
        dialog.setTitle("Transaction Hash Input Dialog");
        dialog.setHeaderText("API: " + apiUrl);
        dialog.setContentText("hash: ");
        dialog.setResizable(true);

        // Traditional way to get the response value.
        Optional<String> result = dialog.showAndWait();
        if (result.isPresent()){
            try {
                openFile(new URL(apiUrl.replace("<hash>", result.get())));
            } catch (MalformedURLException ignored) {}
        }
    }

    private void showFileChooser(FileType ft) {
        File file = MyFileChooser.showFileChooser(stage, ft);
        if (file != null) {
            openFile(file);
        }
    }

    private void openFile(File file) {
        try {
            openFile(file.toURI().toURL());
        } catch (MalformedURLException e) {
            e.printStackTrace(System.err);
        }
    }

    private void openFile(URL url) {
        Tab tab = createTab(url);
        OpenFileTask task = new OpenFileTask(url);

        task.setOnSucceeded((OpenFileResult ofr) -> {
            if (ofr.fileType == FileType.JAVA_JAR) {
                JarTreeView treeView = new JarTreeView(ofr.url, ofr.jarRootNode);
                treeView.setOpenClassHandler(this::openClassInJar);
                tab.setContent(treeView.getTreeView());
            } else {
                ParsedViewerPane viewerPane = new ParsedViewerPane(ofr.fileRootNode, ofr.hexText);
                tab.setContent(viewerPane);
            }

            RecentFiles.INSTANCE.add(ofr.fileType, url);
            menuBar.updateRecentFiles();
        });

        task.setOnFailed((Throwable err) -> {
            Text errMsg = new Text(err.toString());
            tab.setContent(errMsg);
        });

        task.startInNewThread();
    }

    private void openClassInJar(String url) {
        try {
            openFile(new URL(url));
        } catch (MalformedURLException e) {
            e.printStackTrace(System.err);
        }
    }


    public static void main(String[] args) {
        Application.launch(args);
    }

}
